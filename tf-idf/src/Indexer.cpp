#include "Indexer.hpp"
#include <iostream>
#include <cmath>

Indexer::Indexer(std::vector<std::string> file_paths) : file_paths(file_paths) {
    printf("CREATING GLOBAL GLOSSARY\n");
    create_global_glossary();

    TOTAL_DOCS = file_paths.size();
    std::cout << "UNIQUE TERMS IN GLOBAL GLOSSARY -> " << global_glossary.size() << std::endl;
    std::cout << "TOTAL DOCUMENTS -> " << TOTAL_DOCS << std::endl;
}

float Indexer::tf(int term_count, int total_terms) {
    float term_freq = 0.f;
    return term_freq = (float)term_count/total_terms;
}

float Indexer::idf(const int TOTAL_DOCS, int doc_freq) {
    float idf = std::log10((float)TOTAL_DOCS/doc_freq); // don't know which log to use
    return idf;
}

std::pair<std::vector<std::string>, std::vector<std::vector<float>>> Indexer::create_index() {
    init_tf_df_index();
    idf_index = compute_idf();

    std::vector<std::string> doc_names;
    std::vector<std::vector<float>> doc_vectors;
    for (int i = 0; i < TOTAL_DOCS; i++) {
        std::unordered_map<std::string, float> doc_tf = term_freq_index[file_paths[i]];
        std::vector<float> doc_vector = create_tf_idf_vector(doc_tf);

        std::cout << "INDEXING -> " << file_paths[i] << std::endl;

        doc_names.push_back(file_paths[i]);
        doc_vectors.push_back(doc_vector);
    }

    return { doc_names, doc_vectors};
}

std::unordered_map<std::string, float> Indexer::create_tf_vector(std::vector<std::string> tokens, int total_terms) {
    std::unordered_map<std::string, float> tf_vector;

    for (const auto term: global_glossary) {
        // TERM FREQUENCY
        int token_freq_doc = token_count(term, tokens); // we're counting in the current document

        float term_freq = tf(token_freq_doc, total_terms);

        tf_vector[term] = term_freq;
    }

    return tf_vector;
}
void Indexer::create_global_glossary()
{
    for (int idx = 0; idx < file_paths.size(); idx++) {
        std::vector<std::string> tokens = create_tokens(read_file(file_paths[idx]));

        global_glossary.insert(tokens.begin(), tokens.end());
    }
}

void Indexer::create_term_freq_index(int idx, std::vector<std::string> tokens, int total_terms) {
    for (const auto term: global_glossary) {
        // TERM FREQUENCY
        int token_freq_doc = token_count(term, tokens); // we're counting in the current document

        float term_freq = tf(token_freq_doc, total_terms);

        term_freq_index[file_paths[idx]][term] = term_freq;
    }
}

void Indexer::create_df_index(int total_terms, std::vector<std::string> tokens) {
    for (int i = 0; i < total_terms; i++) {
        // DOCUMENT FREQUENCY
        int doc_freq = doc_freq_index[tokens[i]]; // will be 0 if key doesn't exist
        doc_freq += 1;
        doc_freq_index[tokens[i]] = std::min(doc_freq, TOTAL_DOCS);
    }
}

std::unordered_map<std::string, float> Indexer::compute_idf() {
    std::unordered_map<std::string, float> idf_index;
    for (const auto& it: doc_freq_index) {
        float inv_df = idf(TOTAL_DOCS, it.second);
        // std::cout << "computed term: " + it.first << " idf: " << inv_df << "\n"; 
        idf_index[it.first] = inv_df;
    }

    return idf_index;
}

std::vector<float> Indexer::create_tf_idf_vector(std::unordered_map<std::string, float> doc_tf) {
    std::vector<float> tf_idf_vec;
    for (const auto& itr: doc_tf) {
        float tf = itr.second;
        float idf = idf_index[itr.first];

        float tf_idf = tf * idf;
        // printf("tf-idf: %f * %f = %f \n", tf, idf, tf_idf);

        tf_idf_vec.push_back(tf_idf);
    }

    return tf_idf_vec;
}

void Indexer::init_tf_df_index() {
    for (int idx = 0; idx < file_paths.size(); idx++) {
        std::vector<std::string> tokens = create_tokens(read_file(file_paths[idx]));

        int total_terms = tokens.size(); // total terms for the curr doc

        create_term_freq_index(idx, tokens, total_terms);
        create_df_index(total_terms, tokens);
    }
}