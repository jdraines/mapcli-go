build:
    go build -o mapcli

mapped-clis-on-path shell="bash":
    echo 'PATH:~/.mapcli/bin' >> ~/.{{shell}}rc

mapcli-on-path path="./mapcli":
    ln -s $(realpath {{path}}) ~/.local/bin/mapcli


