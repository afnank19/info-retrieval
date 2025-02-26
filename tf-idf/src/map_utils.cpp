#include "../include/map_utils.hpp"

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