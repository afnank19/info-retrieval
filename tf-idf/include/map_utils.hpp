#ifndef MAP_UTILS_H
#define MAP_UTILS_H

#include <unordered_map>
#include <string>
#include <iostream>
#include <iomanip>

void print_term_freq_index(const std::unordered_map<std::string, std::unordered_map<std::string, float>>& table);
void print_string_int_map(const std::unordered_map<std::string, int>& myMap);
void print_string_float_map(const std::unordered_map<std::string, float>& myMap);

#endif