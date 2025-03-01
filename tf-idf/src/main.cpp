#include <iostream>
#include <string>
#include <vector>

#include "file_utils.hpp"
#include "Indexer.hpp"
#include "QueryRunner.hpp"
#include "crow.h"

// Should probably move this elsewhere, but since main is not that complex, it is fine here
std::vector<std::pair<std::string, float>> execute_query(std::string query, Indexer indexer, std::pair<std::vector<std::string>, std::vector<std::vector<float>>> documents) {
    std::vector<std::string> query_tokens = query_tokenizer(query);

    auto tf_vec = indexer.create_tf_vector(query_tokens, query_tokens.size());
    std::vector<float> query_tf_idf_vec = indexer.create_tf_idf_vector(tf_vec);

    auto search_result = QueryRunner::run_query(query_tf_idf_vec, documents);

    return search_result;
}

int main() {
    std::cout << "----BEGIN-ENGINE----" << std::endl;

    std::vector<std::string> file_paths = read_files_from_dir("./collection");

    Indexer indexer(file_paths);
    auto documents = indexer.create_index();

    for ( int i = 0; i < documents.first.size(); i++) {
         std::cout << documents.first[i] << " - vec: [";
        for (int j = 0; j < documents.second[i].size(); j++) {
            std::cout << documents.second[i][j] << ",";
        }
        std::cout << "] \n";
    }

    // std::string query;
    // std::cout << "ENTER QUERY: ";

    crow::SimpleApp app;
    CROW_ROUTE(app, "/search")
    ([indexer, documents](const crow::request& req){
        auto query = crow::query_string{req.url_params};
        std::string input = query.get("q") ? query.get("q") : "";

        auto search_result = execute_query(input, indexer, documents);

        crow::json::wvalue res;
        for (const auto& itr: search_result) {
            std::cout << itr.first << " rel -> " << itr.second << "\n";
            res[itr.first] = itr.second;
        }

        return res;
    });

    CROW_ROUTE(app, "/")
    ([] {
        return "engine-v0 up";
    });

    app.port(5173).multithreaded().run();

    return 0;
}