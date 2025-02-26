#include <iostream>
#include <fstream>
#include <string>
#include <filesystem>
#include <vector>
#include <unordered_set>
#include <unordered_map>
#include <algorithm>
#include <cmath>
#include "../include/file_utils.hpp"
#include "../include/map_utils.hpp"

namespace fs = std::filesystem;

struct Document {
    std::string name;
    std::vector<float> vector;
};

int token_count(std::string token ,std::vector<std::string> tokens) {
    int count = 0;
    for (int i = 0; i < tokens.size(); i++) {
        if (token == tokens[i]) {
            count++;
        }
    }

    return count;
}

float tf(int term_count, int total_terms ) {
    float term_freq = 0.f;
    return term_freq = (float)term_count/total_terms;
}

float idf(const int TOTAL_DOCS, int doc_freq) {
    float idf = std::log10((float)TOTAL_DOCS/doc_freq); // don't know which log to use
    return idf;
}

std::unordered_map<std::string, float> compute_idf(std::unordered_map<std::string, int> doc_freq_index, const int TOTAL_DOCS) {
    std::unordered_map<std::string, float> idf_index;
    for (const auto& it: doc_freq_index) {
        float inv_df = idf(TOTAL_DOCS, it.second);
        // std::cout << "computed term: " + it.first << " idf: " << inv_df << "\n"; 
        idf_index[it.first] = inv_df;
    }

    return idf_index;
}

std::vector<float> create_tf_idf_vector(std::unordered_map<std::string, float> doc_tf, std::unordered_map<std::string, float> idf_index) {
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

int main() {
    std::cout << "Hello, TF-IDF" << std::endl;
    std::vector<std::string> file_paths = read_files_from_dir("./collection");
    // read_file("./collection/doc1.txt"); // path relative to the binary
    std::unordered_map<std::string, std::unordered_map<std::string, float>> term_freq_index;
    std::unordered_map<std::string, float> idf_index;
    std::unordered_map<std::string, int> doc_freq_index;

    std::unordered_set<std::string> global_glossary;

    for (int idx = 0; idx < file_paths.size(); idx++) {
        std::vector<std::string> tokens = create_tokens(read_file(file_paths[idx]));

        global_glossary.insert(tokens.begin(), tokens.end());
    }

    // Creates the TF table
    const int TOTAL_DOCS = file_paths.size(); 
    for (int idx = 0; idx < file_paths.size(); idx++) {
        std::vector<std::string> tokens = create_tokens(read_file(file_paths[idx]));

        int total_terms = tokens.size(); // total terms for the curr doc

        for (const auto term: global_glossary) {
            // TERM FREQUENCY
            int token_freq_doc = token_count(term, tokens); // we're counting in the current document

            float term_freq = tf(token_freq_doc, total_terms);

            // printf("tf: %f,  tfd: %d, tt: %d ",term_freq, token_freq_doc, total_terms);
            // std::cout << tokens[i] << "\n";

            term_freq_index[file_paths[idx]][term] = term_freq;

            // DOCUMENT FREQUENCY
            // int doc_freq = doc_freq_index[term]; // will be 0 if key doesn't exist
            // doc_freq += 1;
            // doc_freq_index[term] = std::min(doc_freq, TOTAL_DOCS);
        }

                // Prev-version
         for (int i = 0; i < total_terms; i++) {
             // DOCUMENT FREQUENCY
             int doc_freq = doc_freq_index[tokens[i]]; // will be 0 if key doesn't exist
             doc_freq += 1;
             doc_freq_index[tokens[i]] = std::min(doc_freq, TOTAL_DOCS);
         }
    }

    print_term_freq_index(term_freq_index);
    // print_string_int_map(doc_freq_index);

    idf_index = compute_idf(doc_freq_index, TOTAL_DOCS);
    print_string_float_map(idf_index);

    // std::unordered_map<std::string, float> doc_tf = term_freq_index["./collection/doc1.txt"];
    // print_string_float_map(doc_tf);

    // std::vector<float> doc1_vec = create_tf_idf_vector(doc_tf, idf_index);

    Document docs[TOTAL_DOCS];

    // Create vectors for each document
    for (int i = 0; i < TOTAL_DOCS; i++) {
        std::unordered_map<std::string, float> doc_tf = term_freq_index[file_paths[i]];
        std::vector<float> doc_vector = create_tf_idf_vector(doc_tf, idf_index);

        docs[i].name = file_paths[i];
        docs[i].vector = doc_vector;
    }

    for (int i = 0; i < TOTAL_DOCS; i++) {
        std::cout << docs[i].name << " - vec: [";
        for (int j = 0; j < docs[i].vector.size(); j++) {
            std::cout << docs[i].vector[j] << ",";
        }
        std::cout << "]" << " size: " << docs[i].vector.size() << "\n";
    }
    return 0;
}



        // Prev-version
        // for (int i = 0; i < total_terms; i++) {
        //     // TERM FREQUENCY
        //     int token_freq_doc = token_count(tokens[i], tokens);

        //     float term_freq = tf(token_freq_doc, total_terms);

        //     // printf("tf: %f,  tfd: %d, tt: %d ",term_freq, token_freq_doc, total_terms);
        //     // std::cout << tokens[i] << "\n";

        //     term_freq_index[file_paths[idx]][tokens[i]] = term_freq;

        //     // DOCUMENT FREQUENCY
        //     int doc_freq = doc_freq_index[tokens[i]]; // will be 0 if key doesn't exist
        //     doc_freq += 1;
        //     doc_freq_index[tokens[i]] = std::min(doc_freq, TOTAL_DOCS);
        // }