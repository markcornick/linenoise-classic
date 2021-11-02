# Installation

=== "Homebrew"
    ```bash
    brew install markcornick/tap/linenoise-classic
    ```

=== "Scoop"
    ```bash
    scoop bucket add markcornick https://github.com/markcornick/scoop-bucket.git
    scoop install linenoise-classic
    ```

=== "GoFish"
    ```bash
    gofish rig add https://github.com/markcornick/gofish-rig.git
    gofish install linenoise-classic
    ```

=== "Snap"
    ```bash
    snap install linenoise-classic
    ```

=== "Docker"
    ```bash
    docker run --rm markcornick/linenoise-classic
    # or
    docker run --rm ghcr.io/markcornick/linenoise-classic
    ```

=== "Binaries and RPM/DEB/APK packages"

    [Here.](https://github.com/markcornick/linenoise-classic/releases)
    Checksums from that page are signed with [my GPG
    key](https://github.com/markcornick.gpg).

    RPM and DEB packages are also available from my Yum and Apt
    repositories, respectively:

    ```bash
    # as root
    cat << E0F > /etc/yum.repos.d/markcornick.repo
    [markcornick]
    name=markcornick repo
    baseurl=https://yum.fury.io/markcornick/
    enabled=1
    gpgcheck=0
    E0F
    ```

    ```bash
    # as root
    echo 'deb [trusted=yes] https://apt.fury.io/markcornick/ /' > /etc/apt/sources.list.d/markcornick.list
    apt-get update
    ```
