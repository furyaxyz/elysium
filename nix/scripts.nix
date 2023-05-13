{ pkgs
, config
, elysium ? (import ../. { inherit pkgs; })
}: rec {
  start-elysium = pkgs.writeShellScriptBin "start-elysium" ''
    # rely on environment to provide elysiumd
    export PATH=${pkgs.test-env}/bin:$PATH
    ${../scripts/start-elysium} ${config.elysium-config} ${config.dotenv} $@
  '';
  start-geth = pkgs.writeShellScriptBin "start-geth" ''
    export PATH=${pkgs.test-env}/bin:${pkgs.go-ethereum}/bin:$PATH
    source ${config.dotenv}
    ${../scripts/start-geth} ${config.geth-genesis} $@
  '';
  start-scripts = pkgs.symlinkJoin {
    name = "start-scripts";
    paths = [ start-elysium start-geth ];
  };
}
