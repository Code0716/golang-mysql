const { defaults } = require('jest-config');

module.exports = {
  roots: ['/web/'],
  testRegex: '((\\.|/)(test|spec))\\.(jsx?|tsx?)$', // テスト対象ファイルの指定
  moduleFileExtensions: [...defaults.moduleFileExtensions, 'ts', 'tsx'],
  moduleNameMapper: { '^@/(.+)': '/web/src/$1' }, // @/はsrc/のalias
  // Setup Enzyme
  snapshotSerializers: ['enzyme-to-json/serializer'],
  setupFilesAfterEnv: ['/web/src/setupEnzyme.ts '],
  transform: { '\\.(js|ts|tsx)?$': 'babel-jest' },
};
