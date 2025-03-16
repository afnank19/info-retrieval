import fs from "fs";
import * as cheerio from "cheerio";

const html = fs.readFileSync('example.html', 'utf8');

// Load into Cheerio
const $ = cheerio.load(html);

// Extract text (removes all tags)
const text = $('body').text().trim();

console.log("Extracted Text:", text);

// TODO: Dir walk and read all html files
// TODO: Parse the HTML and turn it into text files and add to collections folder