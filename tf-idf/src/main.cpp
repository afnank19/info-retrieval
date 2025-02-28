#include <iostream>
#include <string>
#include <vector>
#include "../include/file_utils.hpp"
#include "../include/map_utils.hpp"
#include "../include/Indexer.hpp"
// #include "../include/VectorMath.hpp"
#include "../include/QueryRunner.hpp"

int main() {
    std::cout << "----BEGIN-ENGINE----" << std::endl;

    std::vector<std::string> file_paths = read_files_from_dir("./collection");

    Indexer indexer(file_paths);
    auto documents = indexer.create_index();

    for ( int i = 0; i < documents.first.size(); i++) {
         std::cout << documents.first[i] << " - vec: [";
        for (int j = 0; j < documents.second[i].size(); j++) {
            std::cout << documents.second[i][j] << ",";
        }
        std::cout << "] \n";
    }

    std::string query;
    std::cout << "ENTER QUERY: ";
    std::getline(std::cin, query);
    std::vector<std::string> query_tokens = query_tokenizer(query);

    auto tf_vec = indexer.create_tf_vector(query_tokens, query_tokens.size());
    std::vector<float> query_tf_idf_vec = indexer.create_tf_idf_vector(tf_vec);

    auto search_result = QueryRunner::run_query(query_tf_idf_vec, documents);

    for (const auto& itr: search_result) {
        std::cout << itr.first << " rel -> " << itr.second << "\n";
    }

    return 0;
}