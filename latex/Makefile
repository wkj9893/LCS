lab.pdf: lab.tex
	xelatex lab.tex
	xelatex lab.tex
	make clean
clean:
	fd -E Makefile -E '*.tex' -E '*.pdf' -E images | xargs rm -f