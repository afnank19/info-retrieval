#include <iostream>
#include <fstream>
#include <string>
#include <filesystem>
#include <vector>
#include <unordered_map>

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

std::vector<std::string> create_tokens(std::string str) {
    std::vector<char> temp;
    std::vector<std::string> tokens;

    for (int i = 0; i < str.length(); ++i) {
        if (str[i] == ' ' || i+1 == str.length()) { // needs better tokenization
            std::string token(temp.begin(), temp.end());
            tokens.push_back(token);
            temp.clear();    
            i++;
        }
        temp.push_back(str[i]);
    }

    return tokens;
}

int token_count(std::string token ,std::vector<std::string> tokens) {
    int count = 0;
    for (int i = 0; i < tokens.size(); i++) {
        if (token == tokens[i]) {
            count++;
        }
    }

    return count;
}

// GPT -- debug func
void print_term_freq_index(const std::unordered_map<std::string, std::unordered_map<std::string, float>>& table) {
    // Find all unique column headers
    std::unordered_map<std::string, bool> columnHeaders;
    for (const auto& row : table) {
        for (const auto& col : row.second) {
            columnHeaders[col.first] = true;
        }
    }

    // Print the header row
    std::cout << std::setw(12) << " " << " | ";
    for (const auto& col : columnHeaders) {
        std::cout << std::setw(10) << col.first << " | ";
    }
    std::cout << "\n" << std::string(12 + columnHeaders.size() * 14, '-') << "\n";

    // Print each row with its values
    for (const auto& row : table) {
        std::cout << std::setw(12) << row.first << " | ";
        for (const auto& col : columnHeaders) {
            if (row.second.find(col.first) != row.second.end()) {
                std::cout << std::setw(10) << row.second.at(col.first) << " | ";
            } else {
                std::cout << std::setw(10) << "-" << " | ";  // Print '-' if no value exists
            }
        }
        std::cout << "\n";
    }
}

int main() {
    std::cout << "Hello, TF-IDF" << std::endl;
    std::vector<std::string> file_paths = read_files_from_dir("./collection");
    // read_file("./collection/doc1.txt"); // path relative to the binary
    std::unordered_map<std::string, std::unordered_map<std::string, float>> term_freq_index;

    // Hard coded case for one document, iterate over all documents and run this
    for (int idx = 0; idx < file_paths.size(); idx++) {
        std::vector<std::string> tokens = create_tokens(read_file(file_paths[idx]));
        int total_terms = tokens.size(); // total terms for the curr doc
        for (int i = 0; i < total_terms; i++) {
            int token_freq_doc = token_count(tokens[i], tokens);

            float term_freq = (float)token_freq_doc/total_terms;
            printf("tf: %f,  tfd: %d, tt: %d ",term_freq, token_freq_doc, total_terms);
            std::cout << tokens[i] << "\n";

            term_freq_index[file_paths[idx]][tokens[i]] = term_freq; 
        }
    }


    print_term_freq_index(term_freq_index);
    // Make a map of map for cross tabulation tables
    return 0;
}


