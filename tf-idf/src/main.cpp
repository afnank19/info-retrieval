#include <iostream>
#include <string>
#include <vector>
#include "../include/file_utils.hpp"
#include "../include/map_utils.hpp"
#include "../include/Indexer.hpp"

int main() {
    std::cout << "Hello, TF-IDF" << std::endl;
    std::vector<std::string> file_paths = read_files_from_dir("./collection");
    // read_file("./collection/doc1.txt"); // path relative to the binary

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
    std::cout << "Enter query: ";
    std::getline(std::cin, query);
    std::cout << "You entered: " << query << "\n";
    std::vector<std::string> query_tokens = query_tokenizer(query);

    for (int i = 0; i < query_tokens.size(); i++) {
        std::cout << query_tokens[i] << ",";
    }

    auto tf_vec = indexer.create_tf_vector(query_tokens, query_tokens.size());

    std::cout << "Term freqs for QUERY are: \n";

    for ( auto& itr: tf_vec) {
        std::cout << "Term: " << itr.first << " tf: " << itr.second << "\n";
    }

    std::vector<float> query_tf_idf_vec = indexer.create_tf_idf_vector(tf_vec);

    std::cout << "TF_IDF vector for query is -> vec: [";
    for (int i = 0; i < query_tf_idf_vec.size(); i++) {
        std::cout << query_tf_idf_vec[i] << ",";
    }
    std::cout << "]";

    return 0;
}