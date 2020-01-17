module.exports = {
  extends: [
    'react-app', // for editor
    'eslint:recommended',
    'plugin:prettier/recommended',
    'prettier/react',
  ],
  rules: {
    quotes: ['error', 'single'],
    'no-debugger': process.env.NODE_ENV === 'production' ? 2 : 0,
    'no-console': process.env.NODE_ENV === 'production' ? 2 : 0,
    'no-unused-expressions': 'off',
    'prettier/prettier': [
      'error',
      {
        singleQuote: true,
        semi: false,
        trailingComma: 'all',
        jsxBracketSameLine: true,
      },
    ],
  },
}
