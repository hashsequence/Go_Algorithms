process.stdin.resume()
process.stdin.setEncoding("ascii")

var input = []

/**
 * [Standard IO collection]
 */
process.stdin.on("data", (chunk) => input.push(chunk.trim()))
process.stdin.on("end", function () { //press ctrl d on ubuntu bash terminal
    input.forEach((docs) => {
    process.stdout.write(docs+'\n')
  })
})
