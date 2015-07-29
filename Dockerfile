FROM golang:1.4.2-onbuild

RUN git config --system user.name "Git Monitor"
RUN git config --system user.email git-monitor@bg.vnet.ibm.com
RUN git config --system credential.helper 'store --file .git-credentials'
