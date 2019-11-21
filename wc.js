const fs = require('fs');
const assert = require('assert');

const filenames = ["readme.md", "wc.js", "other.txt"];

for (file of filenames) {
  // const file = filename;
  console.log("Reading File:", file);
  if (!fs.existsSync(file)) {
    console.log("File does not exist:", file);
    continue;
  }
  const readStream = fs.createReadStream(file, {autoClose: false, emitClose: true});
  readStream.on('open', () => {
    console.log("opened file", file);
    readStream.on("data", (chunk) => {
      assert.equal(done, false);
      count += chunk.length;
    });
  });

  // Handling destruction of process
  let count = 0;
  let done = false;
  readStream.on("end", () => {
    done = true;
    console.log(file, count);
    // readStream.destroy();
  });
  readStream.on('close', () => {
    console.log('closed');
  });
}

// Do we need this setTimeout
// setTimeout(() => {}, 1000);
