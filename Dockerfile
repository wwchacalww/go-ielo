FROM golang:1.20rc1-bullseye

RUN useradd -ms /bin/bash chacal


WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN  go install github.com/golang/mock/mockgen@v1.6.0 && \
  go install github.com/spf13/cobra-cli@latest

RUN mkdir -p /usr/share/man/man1 && \
  echo 'deb http://ftp.debian.org/debian stretch-backports main' | tee /etc/apt/sources.list.d/stretch-backports.list && \
  apt update -y && apt install -y \
  sqlite3 \
  git \
  ca-certificates \
  zsh \
  curl \
  wget \
  fonts-powerline \
  procps


RUN usermod -u 1000 chacal
RUN mkdir -p /var/www/.cache
RUN chown -R chacal:chacal /go
RUN chown -R chacal:chacal /var/www/.cache
USER chacal

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.2/zsh-in-docker.sh)" -- \
  -t https://github.com/romkatv/powerlevel10k \
  -p git \
  -p git-flow \
  -p https://github.com/zdharma-continuum/fast-syntax-highlighting \
  -p https://github.com/zsh-users/zsh-autosuggestions \
  -p https://github.com/zsh-users/zsh-completions \
  -a 'export TERM=xterm-256color'

RUN echo '[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh' >> ~/.zshrc && \
  echo 'HISTFILE=/go/src/.zsh/.zsh_history' >> ~/.zshrc

CMD ["tail", "-f", "/dev/null"]