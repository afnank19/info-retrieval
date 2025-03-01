#include "utils.hpp"

int token_count(std::string token ,std::vector<std::string> tokens) {
    int count = 0;
    for (int i = 0; i < tokens.size(); i++) {
        if (token == tokens[i]) {
            count++;
        }
    }

    return count;
}