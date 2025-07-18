local M = {}

function M.is_windows()
    return vim.loop.os_uname().version:match("Windows")
end

function M.file_exists(fname)
    local stat = vim.loop.fs_stat(fname)
    return (stat and stat.type) or false
end

function M.is_git_repo()
    local git_dir = vim.fn.finddir(".git", ".;")
    return git_dir ~= ""
end

function M.join_paths(...)
    local path_sep = M.is_windows() and "\\" or "/"
    local result = table.concat({ ... }, path_sep)
    return result
end

function M.get_config_dir()
    local config = vim.fn.stdpath("config")
    return config
end

function M.reload_module(name)
    require("plenary.reload").reload_module(name)
end

function M.inspect(v)
    print(vim.inspect(v))
    return v
end

function M.is_nvim_0_9()
    local actual = vim.version()
    return actual and (actual.major == 0 and actual.minor >= 9)
end

function M.is_compatible()
    local actual = vim.version()
    local required = { major = 0, minor = 9 }

    if not actual or not actual.major or not actual.minor then
        vim.notify("Unable to detect Neovim version properly!", vim.log.levels.ERROR)
        return false
    end

    for k, v in pairs(required) do
        if (actual[k] or 0) < v then
            return false
        end
    end
    return true
end

return M
