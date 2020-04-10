const rootDir = `${__dirname}/dist/`;

module.exports = [
  {
    // メインとなるJavaScriptファイル（エントリーポイント）
    entry: `./src/index.js`,

    // ファイルの出力設定
    output: {
      //  出力ファイルのディレクトリ名
      path: rootDir + `/js/`,
      // 出力ファイル名
      filename: 'main.js',
    },
    module: {
      rules: [
        {
          // 拡張子 .js の場合
          test: /\.js$/,
          use: [
            {
              // Babel を利用する
              loader: 'babel-loader',
              // Babel のオプションを指定する
              options: {
                presets: [
                  ['@babel/react'],
                  [
                    '@babel/env',
                    {
                      targets: {
                        ie: 11,
                      },
                      useBuiltIns: 'usage',
                      corejs: 3,
                    },
                  ],
                ],
                plugins: ['@babel/plugin-proposal-class-properties'],
              },
            },
          ],
          // node_modules は除外する
          exclude: /node_modules/,
        },
        {
          test: /\.(scss)$/,
          use: [
            {
              loader: 'style-loader', // inject CSS to page
            },
            {
              loader: 'css-loader', // translates CSS into CommonJS modules
            },
            {
              loader: 'postcss-loader', // Run post css actions
              options: {
                plugins: () => {
                  // post css plugins, can be exported to postcss.config.js
                  return [require('precss'), require('autoprefixer')];
                },
              },
            },
            {
              loader: 'sass-loader', // compiles Sass to CSS
            },
          ],
        },
        {
          test: /\.(css)$/,
          loaders: [
            'style-loader',
            'css-loader?modules&importLoaders=1&localIdentName=[name]__[local]___[hash:base64:5]',
          ],
          exclude: [/node_modules/],
        },
        {
          test: /\.(css)$/,
          loaders: ['style-loader', 'css-loader'],
          include: [/node_modules/],
        },
      ],
    },
    //ソースマップはbuild時削除
    devtool: 'inline-source-map',
    // performance: { hints: false },
    // ローカル開発用環境を立ち上げる
    // 実行時にブラウザが自動的に localhost を開く
    devServer: {
      contentBase: rootDir,
      publicPath: '/js/',
      disableHostCheck: true,
      open: true,
      watchContentBase: true,
      port: 8080,
      host: '0.0.0.0',
    },
  },
];
