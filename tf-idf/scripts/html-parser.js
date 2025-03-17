import fs from 'fs';
import path from 'path';
import * as cheerio from "cheerio";
import { fileURLToPath } from 'url';

// Derive __dirname since it's not available in ES modules.
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Define the collection directory relative to this script's location.
const collectionDir = path.resolve(__dirname, '../collection');

// Create the collection directory if it doesn't exist.
if (!fs.existsSync(collectionDir)) {
  fs.mkdirSync(collectionDir, { recursive: true });
  console.log(`Created collection directory: ${collectionDir}`);
}

/**
 * Recursively walks through directories starting at "dir", ignoring hidden folders.
 * For each HTML file found, it calls processHtmlFile().
 */
async function walkDir(dir) {
  let entries;
  try {
    entries = await fs.promises.readdir(dir, { withFileTypes: true });
  } catch (err) {
    console.error(`Error reading directory ${dir}: ${err}`);
    return;
  }

  for (const entry of entries) {
    // Skip hidden files and folders (names starting with a dot)
    if (entry.name.startsWith('.')) continue;

    const fullPath = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      await walkDir(fullPath);
    } else if (entry.isFile() && /\.html?$/i.test(entry.name)) {
      await processHtmlFile(fullPath);
    }
  }
}

/**
 * Reads an HTML file, extracts its text using Cheerio, and writes the text
 * to a new .txt file in the collection directory.
 */
async function processHtmlFile(filePath) {
  try {
    const html = await fs.promises.readFile(filePath, 'utf8');
    const $ = cheerio.load(html);
    const text = $.text();

    // Create a safe file name based on the full path.
    // Replace characters not allowed in file names with an underscore.
    const safeName = filePath.replace(/[/\\?%*:|"<>]/g, '_');
    const outputPath = path.join(collectionDir, safeName + '.txt');

    await fs.promises.writeFile(outputPath, text);
    console.log(`Processed: ${filePath} -> ${outputPath}`);
  } catch (err) {
    console.error(`Error processing file ${filePath}: ${err}`);
  }
}

// Start the process from the desired root directory.
// Change process.cwd() to any starting directory if needed.
(async () => {
  try {
    const startDir = process.cwd();
    console.log(`Starting directory traversal from: ${startDir}`);
    await walkDir(startDir);
    console.log('Processing complete.');
  } catch (err) {
    console.error(`Error during processing: ${err}`);
  }
})();
