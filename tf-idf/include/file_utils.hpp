#ifndef FILE_UTILS_H
#define FILE_UTILS_H

#include <string>
#include <vector>

std::string read_file(std::string filename);
std::vector<std::string> read_files_from_dir(std::string path);
std::vector<std::string> create_tokens(std::string str);
std::vector<std::string> query_tokenizer(std::string str);

#endif