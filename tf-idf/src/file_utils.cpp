#include "file_utils.hpp"
#include <iostream>
#include <fstream>
#include <filesystem>

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
        if (curr_line.length() > 13) 
        {
            std::string flag = curr_line.substr(0, 13);

            if ( flag == "<title_swe12>") {
                curr_line = curr_line.substr(13, curr_line.length());
            }

            if ( flag == "<link_swe123>") {
                curr_line = curr_line.substr(13, curr_line.length());
            }
        }

        file_text += curr_line + " "; // new line has no spaces
    }

    // std::cout << "final str: " + file_text + "\n";
    curr_file.close();

    return file_text;
}

// Horribly written
std::pair<std::string, std::string> get_title_and_link_from_file(std::string filename)
{
    std::ifstream curr_file(filename);
    std::string title, link, curr_line;

    if( curr_file.fail( ) ) {
            std::cerr << "Error - Failed to open " << filename << std::endl;
            exit( -1 );  // Or use a loop to ask for a different file name.
    }

    while (getline(curr_file, curr_line)) {
        std::string flag = curr_line.substr(0, 13);
        bool foundAnyFlag = false;

        if ( flag == "<title_swe12>") {
            title = curr_line.substr(13, curr_line.length());
            foundAnyFlag = true;
        }

        if ( flag == "<link_swe123>") {
            link = curr_line.substr(13, curr_line.length());
            foundAnyFlag = true;
        }

        if (!foundAnyFlag) {
            break;
        }
    }
    curr_file.close();

    return { title, link };
}

std::vector<std::string> read_files_from_dir(std::string path) {
    std::vector<std::string> file_paths;

    try {
        for (const auto& entry : fs::directory_iterator(path)) {
            // Check if the entry is a regular file (optional)
            if (fs::is_regular_file(entry.status())){
                std::cout <<"Adding: " << entry.path() << std::endl;
                file_paths.push_back(entry.path());
            }
        }
    }
    catch(const std::exception& e){
        std::cerr << e.what() << '\n';
    }

    return file_paths;
}
/* JANK ALERT AAAAAAAA*/
// Not really a fan of this being here, but currently no where else to put it
std::vector<std::string> create_tokens(std::string str) {
    std::vector<char> temp;
    std::vector<std::string> tokens;

    for (int i = 0; i < str.length(); ++i) {
        if (str[i] == ' ' || str[i] == '.' || str[i] == ',' || i+1 == str.length()) { // needs better tokenization
            std::string token(temp.begin(), temp.end());
            Porter2Stemmer::trim(token);
            Porter2Stemmer::stem(token);
            tokens.push_back(token);
            // std::cout << token << std::endl;
            temp.clear();    
            i++;
        }
        temp.push_back(std::tolower(str[i]));
    }

    return tokens;
}

// Didnt want to waste time on the tokenizer, this works user input
// the above one is working for txt files, i'll work on the tokenizer later
// as it is not the biggest concern here
std::vector<std::string> query_tokenizer(std::string str) {
    std::vector<char> temp;
    std::vector<std::string> tokens;

    for (int i = 0; i <= str.length(); ++i) {
        if (str[i] == ' ' || i == str.length()) { // needs better tokenization
            std::string token(temp.begin(), temp.end());
            Porter2Stemmer::trim(token);
            Porter2Stemmer::stem(token);
            tokens.push_back(token);
            temp.clear();    
        } else {
            temp.push_back(std::tolower(str[i]));
        }
    }

    return tokens;
}