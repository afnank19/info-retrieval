#include <iostream>
#include <string>
#include <vector>

#include "file_utils.hpp"
#include "Indexer.hpp"
#include "QueryRunner.hpp"
#include "crow.h"

struct Result
{
    std::string title;
    std::string link;
    float relevancy;
};


// Should probably move this elsewhere, but since main is not that complex, it is fine here
std::vector<Result> execute_query(std::string query, Indexer indexer, std::pair<std::vector<std::string>, std::vector<std::vector<float>>> documents) {
    std::vector<std::string> query_tokens = query_tokenizer(query);

    auto tf_vec = indexer.create_tf_vector(query_tokens, query_tokens.size());
    std::vector<float> query_tf_idf_vec = indexer.create_tf_idf_vector(tf_vec);

    auto search_result = QueryRunner::run_query(query_tf_idf_vec, documents); // contains a pair of filepaths and their relevancy scores
    // Need to build the correct result here
    auto res_map = indexer.get_result_map();
    std::vector<std::pair<std::string, std::string>> res;
    std::vector<Result> result;
    for (const auto& itr: search_result) {
            // res[itr.first] = itr.second;
            std::string doc_path = itr.first;
            std::cout << itr.first << "->" << itr.second << "\n";
            std::string title = res_map[doc_path].title;
            std::string link = res_map[doc_path].link;
            res.push_back(std::make_pair(title, link));

            Result r;
            r.title = title;
            r.link = link;
            r.relevancy = itr.second;

            result.push_back(r);
    }


    //std::string test = search_result[1].first;

    //std::cout << res_map[test].title;

    return result;
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
        std::size_t i = 0;
        for (const auto& itr: search_result) {
            //std::cout << itr.first << " rel -> " << itr.second << "\n";

            res[i]["title"] = itr.title;
            res[i]["link"] = itr.link;
            res[i]["relevancy"] = itr.relevancy;
            ++i;
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