# install
npm install -g resume-cli

# dependencies
npm install jsonresume-theme-stackoverflow

# export
hackmyresume build resume.json TO out/resume.html -t node_modules/jsonresume-theme-stackoverflow

resume export resume.html -t spartan

cp out/resume.html  ~/Git/blog/content/