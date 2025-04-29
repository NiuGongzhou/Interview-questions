module.exports = {
    preset: 'ts-jest',
    testEnvironment: 'jsdom',
    globals: {
      'ts-jest': {
        tsconfig: 'tsconfig.test.json'
      }
    },
    moduleNameMapper: {
      '^@/(.*)$': '<rootDir>/src/$1',
      '^@gen/(.*)$': '<rootDir>/src/gen/$1'
    }
  }