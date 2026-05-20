-- You can add your own plugins here or in other files in this directory!
--  I promise not to create any merge conflicts in this directory :)
--
-- See the kickstart.nvim README for more information
return {
  -- 'sphamba/smear-cursor.nvim',
  -- opts = {
  --   stiffness = 0.5,
  --   trailing_stiffness = 0.5,
  --   matrix_pixel_threshold = 0.5,
  -- },
  'folke/snacks.nvim',
  priority = 1000,
  lazy = false,
  opts = {
    image = {
      enabled = true,
      -- force ghostty detection if auto-detect fails:
      -- set env var SNACKS_GHOSTTY=true in your shell or:
    },
  },
}
