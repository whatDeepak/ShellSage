" ShellSage for Vim.
" Set `g:ShellSageBin` path below.
" Mapped to Ctrl+K (you can change bellow).
" On Normal Mode passes the current line.
" Or select many lines.

let g:ShellSageBin = '_'
let g:gretting = '...waiting for ShellSage...'

let g:selectionCallCalled = 0
autocmd CursorMoved * let g:selectionCallCalled = 0

nnoremap <C-k> :<C-u>call SingleLineCall()<CR>
xnoremap <C-k> :<C-u>call SelectionCall()<CR>

function! EscapeSpecialChars(str)
	let escapedStr = substitute(a:str, "'", "''", 'g')
	let escapedStr = substitute(escapedStr, '"', '\\"', 'g')
	let escapedStr = substitute(escapedStr, "`", "\\`", 'g')
	return escapedStr
endfunction

function! CleanAnsi(str)
	let clean = substitute(a:str, '[\033\e\x1B]\?\[[0-9;]*m', '', 'g')
	let clean = substitute(clean, '^>> \+', '', 'g')
	return clean
endfunction

fun! ShellSage(s)
"	echo "Original string: " . a:s
	let s = EscapeSpecialChars(a:s)
"	echo "Escaped string: " . s
	let return_from_ShellSage = system(g:ShellSageBin . " \'" . s . "\'")
	let return_from_ShellSage = trim(return_from_ShellSage)
	let lines = split(return_from_ShellSage, '\v\n')
	let out = []
	for line in lines
		call add(out, CleanAnsi(line))
	endfor
	return out
endfunction

function! PrintOut(out)
	let current_line = line('.')
	call append(current_line, '<<<<<<<')
	call append(current_line, a:out)
	call append(current_line, '>>>>>>>')
endfunction

function! SelectionCall()
	if g:selectionCallCalled == 1
		return
	endif
	g:selectionCallCalled = 1
	echo g:gretting
	let start_pos = getpos("'<")
	let end_pos = getpos("'>")
	let start_line = start_pos[1]
	let end_line = end_pos[1]
	execute "normal " . end_line . "G"
	let lines = getline(start_line, end_line)
	let line = join(lines, '\n')
	let out = ShellSage(line)
	call PrintOut(out)
endfunction

function! SingleLineCall()
	echo g:gretting
	let line = getline('.')
	let out = ShellSage(line)
	call PrintOut(out)
endfunction
