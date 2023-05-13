{ pkgs ? import ../../nix { } }:
let elysiumd = (pkgs.callPackage ../../. { });
in
elysiumd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-elysiumd.patch
  ];
})
