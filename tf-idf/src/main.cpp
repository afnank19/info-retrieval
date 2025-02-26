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
    return 0;
}