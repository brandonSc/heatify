FROM golang:1.4.2-onbuild

RUN git config --system user.name "Brandon Schurman"
RUN git config --system user.email schurman@ca.ibm.com
RUN git config --system credential.helper 'store --file .git-credentials'
