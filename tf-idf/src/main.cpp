#include <iostream>
#include <string>
#include <vector>

#include "file_utils.hpp"
#include "Indexer.hpp"
#include "QueryRunner.hpp"

// Should probably move this elsewhere, but since main is not that complex, it is fine here
std::vector<std::pair<std::string, float>> execute_query(std::string query, Indexer indexer, std::pair<std::vector<std::string>, std::vector<std::vector<float>>> documents) {
    std::vector<std::string> query_tokens = query_tokenizer(query);

    auto tf_vec = indexer.create_tf_vector(query_tokens, query_tokens.size());
    std::vector<float> query_tf_idf_vec = indexer.create_tf_idf_vector(tf_vec);

    auto search_result = QueryRunner::run_query(query_tf_idf_vec, documents);

    return search_result;
}

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

    auto search_result = execute_query(query, indexer, documents);

    for (const auto& itr: search_result) {
        std::cout << itr.first << " rel -> " << itr.second << "\n";
    }

    return 0;
}