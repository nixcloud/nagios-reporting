with import <nixpkgs> { };
with pkgs;

rec {
  gocraft-web = buildGoPackage rec {
    rev = "934a096ff6918ed5d4667075794a8d18a88c1576";
    name = "gocraft-web-${stdenv.lib.strings.substring 0 7 rev}";
    goPackagePath = "github.com/gocraft/web";
    #doCheck = true;

    src = fetchFromGitHub {
      inherit rev;
      owner = "gocraft";
      repo = "web";
      sha256 = "0f09jnhvc8yw2r2xk57qbjaps9qfxs30jgyvyrhjy5xm8l2dvzwx";
    };
  };
}
