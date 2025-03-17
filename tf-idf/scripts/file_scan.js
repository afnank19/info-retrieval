import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

// Define __filename and __dirname for ES modules.
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Determine starting directory based on platform.
const startDir = process.platform === 'win32' ? 'C:\\' : '/';

// Output file: stored in the parent directory of the script.
const outputFilePath = path.join(__dirname, '../html_files.txt');
const output = fs.createWriteStream(outputFilePath, { flags: 'w' });

/**
 * Recursively walk through directories and process .html files.
 * @param {string} dir - The directory to scan.
 */
async function walkDir(dir) {
  let entries;
  try {
    // Get directory entries with file types.
    entries = await fs.promises.readdir(dir, { withFileTypes: true });
  } catch (err) {
    // Skip directories that cannot be read (e.g., due to permissions).
    console.error(`Error reading directory ${dir}: ${err.message}`);
    return;
  }

  for (const entry of entries) {
    const fullPath = path.join(dir, entry.name);

    if (entry.isDirectory()) {
      // Recursively scan subdirectories.
      await walkDir(fullPath);
    } else if (
      entry.isFile() &&
      path.extname(entry.name).toLowerCase() === '.html'
    ) {
      // Write the full path to the output file followed by a newline.
      output.write(fullPath + "\n");
    }
  }
}

(async () => {
  console.log(`Scanning file system starting at ${startDir}...`);
  await walkDir(startDir);
  output.end(() => {
    console.log(`Scanning complete. HTML file paths saved in ${outputFilePath}`);
  });
})();
