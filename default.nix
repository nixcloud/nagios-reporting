with import <nixpkgs> { };

let 
  myDeps =  (import ./myDeps.nix);
in

buildGoPackage rec {
  name = "nagios-reporting-${version}";
  version = "0.0.1";

  src = ./.;

  goPackagePath = "github.com/nixcloud/nagios-reporting";

  buildInputs = with myDeps; [ gocraft-web ]; 
}

