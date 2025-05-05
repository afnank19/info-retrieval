#ifndef INDEXER_H
#define INDEXER_H

#include <unordered_map>
#include <unordered_set>
#include <string>
#include <vector>
#include <utility>
#include "file_utils.hpp"
#include "map_utils.hpp"
#include "utils.hpp"

class Indexer {
private:
    struct Result {
        std::string title;
        std::string link;
    };
    std::unordered_map<std::string, std::unordered_map<std::string, float>> term_freq_index;
    std::unordered_map<std::string, float> idf_index;
    std::unordered_map<std::string, int> doc_freq_index;
    std::unordered_map<std::string, Result> result_map;
    std::unordered_set<std::string> global_glossary;
    std::vector<std::string> file_paths;

    int TOTAL_DOCS;

    void create_global_glossary();
    void create_term_freq_index(int idx, std::vector<std::string> tokens, int total_terms);
    void create_df_index(int total_terms, std::vector<std::string> tokens);
    void init_tf_df_index();
    void build_result_map();
    void print_result_map();
    std::unordered_map<std::string, float> compute_idf();
public:
    Indexer(std::vector<std::string> file_paths);
    std::pair<std::vector<std::string>, std::vector<std::vector<float>>> create_index();
    std::unordered_map<std::string, float> create_tf_vector(std::vector<std::string> tokens, int total_terms);
    std::vector<float> create_tf_idf_vector(std::unordered_map<std::string, float> doc_tf);
    float tf(int term_count, int total_terms );
    float idf(const int TOTAL_DOCS, int doc_freq);

    std::unordered_map<std::string, Result> get_result_map();
    // ~Indexer();
};


#endif