module.exports = {
  module: {
    rules: [
      {
        test: '\.pug$',
        loader: 'pug-plain-loader'
      },
      {
        test: /\.scss$/,
        loaders: ["style", "css", "sass"]
      }
    ]
  }
}
