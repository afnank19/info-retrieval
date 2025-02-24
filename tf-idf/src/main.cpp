#include <iostream>
#include <fstream>
#include <string>
#include <filesystem>
#include <vector>
#include <unordered_map>
#include <algorithm>
#include <cmath>

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

float tf(int term_count, int total_terms ) {
    float term_freq = 0.f;
    return term_freq = (float)term_count/total_terms;
}

// GPT -- debug/utils func
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
void print_string_int_map(const std::unordered_map<std::string, int>& myMap) {
    std::cout << std::setw(15) << "Key" << " | " << std::setw(10) << "Value" << "\n";
    std::cout << std::string(30, '-') << "\n";

    for (const auto& pair : myMap) {
        std::cout << std::setw(15) << pair.first << " | " << std::setw(10) << pair.second << "\n";
    }
}
void print_string_float_map(const std::unordered_map<std::string, float>& myMap) {
    std::cout << std::setw(15) << "Key" << " | " << std::setw(10) << "Value" << "\n";
    std::cout << std::string(30, '-') << "\n";

    for (const auto& pair : myMap) {
        std::cout << std::setw(15) << pair.first << " | " << std::setw(10) << pair.second << "\n";
    }
}

float idf(const int TOTAL_DOCS, int doc_freq) {
    float idf = std::log((float)TOTAL_DOCS/doc_freq); // don't know which log to use
    return idf;
}

std::unordered_map<std::string, float> compute_idf(std::unordered_map<std::string, int> doc_freq_index, const int TOTAL_DOCS) {
    std::unordered_map<std::string, float> idf_index;
    for (const auto& it: doc_freq_index) {
        float inv_df = idf(TOTAL_DOCS, it.second);
        std::cout << "computed term: " + it.first << " idf: " << inv_df << "\n"; 
        idf_index[it.first] = inv_df;
    }

    return idf_index;
}


int main() {
    std::cout << "Hello, TF-IDF" << std::endl;
    std::vector<std::string> file_paths = read_files_from_dir("./collection");
    // read_file("./collection/doc1.txt"); // path relative to the binary
    std::unordered_map<std::string, std::unordered_map<std::string, float>> term_freq_index;
    std::unordered_map<std::string, float> idf_index;
    std::unordered_map<std::string, int> doc_freq_index;

    // Creates the TF table
    const int TOTAL_DOCS = file_paths.size(); 
    for (int idx = 0; idx < file_paths.size(); idx++) {
        std::vector<std::string> tokens = create_tokens(read_file(file_paths[idx]));

        int total_terms = tokens.size(); // total terms for the curr doc

        for (int i = 0; i < total_terms; i++) {
            // TERM FREQUENCY
            int token_freq_doc = token_count(tokens[i], tokens);

            float term_freq = tf(token_freq_doc, total_terms);

            printf("tf: %f,  tfd: %d, tt: %d ",term_freq, token_freq_doc, total_terms);
            std::cout << tokens[i] << "\n";

            term_freq_index[file_paths[idx]][tokens[i]] = term_freq;

            // DOCUMENT FREQUENCY
            int doc_freq = doc_freq_index[tokens[i]]; // will be 0 if key doesn't exist
            doc_freq += 1;
            doc_freq_index[tokens[i]] = std::min(doc_freq, TOTAL_DOCS);
        }
    }

    print_term_freq_index(term_freq_index);
    print_string_int_map(doc_freq_index);

    idf_index = compute_idf(doc_freq_index, TOTAL_DOCS);
    print_string_float_map(idf_index);

    return 0;
}


