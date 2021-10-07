const HTMLWebpackPlugin = require('html-webpack-plugin');
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const path = require("path");
const {BundleAnalyzerPlugin} = require('webpack-bundle-analyzer');

module.exports = {
    entry: "./index.js",
    mode: "development",
    output: {
        filename: "bundle.js",
        path: path.resolve(__dirname, 'dist'),
    },
    plugins: [
        new HTMLWebpackPlugin({
            template: "./index.html"
        }),
        new CleanWebpackPlugin(),
        new BundleAnalyzerPlugin(),
    ],
    module: {
        rules: [
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader']
            },
            {
                test: /\.png$/,
                use: ['file-loader']
            },
            {
                test: /\.jpg$/,
                use: ['file-loader']
            },
        ]
    }
}