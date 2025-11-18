local ls = require 'luasnip'
local s = ls.snippet
local t = ls.text_node

return {
  s('blp', {
    t {
      '#include <stdio.h>',
      '',
      'int main() {',
      '    printf("Hello, World!\\n");',
      '',
      '    return 0;',
      '}',
    },
  }),
}
