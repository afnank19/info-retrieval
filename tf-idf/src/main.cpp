#include <iostream>
#include <fstream>
#include <string>
#include <filesystem>
#include <vector>

namespace fs = std::filesystem;

std::string read_file(std::string filename) {
    std::ifstream curr_file(filename);
    std::string file_text, curr_line;

    if( curr_file.fail( ) ) {
            std::cerr << "Error - Failed to open " << filename << std::endl;
            exit( -1 );  // Or use a loop to ask for a different file name.
    }

    while (getline(curr_file, curr_line)) {
        // std::cout << "READING: " + curr_line + "\n";
        file_text += curr_line + " "; // new line has no spaces
    }

    // std::cout << "final str: " + file_text + "\n";
    curr_file.close();

    return file_text;
}

std::vector<std::string> read_files_from_dir(std::string path) {
    std::vector<std::string> file_paths;

    try {
        for (const auto& entry : fs::directory_iterator(path)) {
            // Check if the entry is a regular file (optional)
            if (fs::is_regular_file(entry.status()))
                std::cout <<"Adding: " << entry.path() << std::endl;
                file_paths.push_back(entry.path());
        }
    }
    catch(const std::exception& e){
        std::cerr << e.what() << '\n';
    }

    return file_paths;
}

int main() {
    std::cout << "Hello, TF-IDF" << std::endl;
    std::vector<std::string> file_paths = read_files_from_dir("./collection");
    // read_file("./collection/doc1.txt"); // path relative to the binary

    for (int i = 0; i < file_paths.size(); i++) {
        std::cout << read_file(file_paths[i]) << std::endl;
        // Read the contents of each file
        // tokenize, and add to some sort of data structure
        // I feel like a vector of a struct would be a decent representation
        // std::vector<Term> term; etc 
        // or unordered_map<string, Term> mp; where Term could be a struct that has a tf, idf, and other metadata needed;
        // and the string can be the term 
    }

    return 0;
}


