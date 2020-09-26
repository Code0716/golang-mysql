module.exports = {
  moduleFileExtensions: ['ts', 'tsx', 'js'],
  transform: {
    '^.+\\.(ts|tsx)$': './node_modules/ts-jest/preprocessor.js',
  },
  moduleNameMapper: {
    '^.*\\.(md|css|scss|jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$':
      '<rootDir>src/testUtils/stub.js',
  },
  setupFilesAfterEnv: ['<rootDir>/src/testUtils/setupTests.ts'],
  moduleDirectories: ['node_modules', 'src'],
};
