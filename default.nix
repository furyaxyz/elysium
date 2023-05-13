{ lib
, stdenv
, buildGoApplication
, nix-gitignore
, buildPackages
, coverage ? false # https://tip.golang.org/doc/go1.20#cover
, rocksdb
, network ? "mainnet"  # mainnet|testnet
, rev ? "dirty"
, static ? stdenv.hostPlatform.isStatic
}:
let
  version = "v1.0.4";
  pname = "elysiumd";
  tags = [ "ledger" "netgo" network "rocksdb" "grocksdb_no_link" ];
  ldflags = lib.concatStringsSep "\n" ([
    "-X github.com/cosmos/cosmos-sdk/version.Name=elysium"
    "-X github.com/cosmos/cosmos-sdk/version.AppName=${pname}"
    "-X github.com/cosmos/cosmos-sdk/version.Version=${version}"
    "-X github.com/cosmos/cosmos-sdk/version.BuildTags=${lib.concatStringsSep "," tags}"
    "-X github.com/cosmos/cosmos-sdk/version.Commit=${rev}"
  ]);
  buildInputs = [ rocksdb ];
in
buildGoApplication rec {
  inherit pname version buildInputs tags ldflags;
  src = (nix-gitignore.gitignoreSourcePure [
    "/*" # ignore all, then add whitelists
    "!/x/"
    "!/app/"
    "!/cmd/"
    "!/client/"
    "!/versiondb/"
    "!/memiavl/"
    "!/store/"
    "!go.mod"
    "!go.sum"
    "!gomod2nix.toml"
  ] ./.);
  modules = ./gomod2nix.toml;
  pwd = src; # needed to support replace
  subPackages = [ "cmd/elysiumd" ];
  buildFlags = lib.optionalString coverage "-cover";
  CGO_ENABLED = "1";
  CGO_LDFLAGS =
    if static then "-lrocksdb -pthread -lstdc++ -ldl -lzstd -lsnappy -llz4 -lbz2 -lz"
    else if stdenv.hostPlatform.isWindows then "-lrocksdb-shared"
    else "-lrocksdb -pthread -lstdc++ -ldl";

  postFixup = lib.optionalString stdenv.isDarwin ''
    ${stdenv.cc.targetPrefix}install_name_tool -change "@rpath/librocksdb.7.dylib" "${rocksdb}/lib/librocksdb.dylib" $out/bin/elysiumd
  '';

  doCheck = false;
  meta = with lib; {
    description = "Official implementation of the Elysium blockchain protocol";
    homepage = "https://elysium.org/";
    license = licenses.asl20;
    mainProgram = "elysiumd" + stdenv.hostPlatform.extensions.executable;
    platforms = platforms.all;
  };
}
