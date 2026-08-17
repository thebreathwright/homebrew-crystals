class Crystals < Formula
  desc "Crystal-native Grax routing and receipt contract assets"
  homepage "https://github.com/ollama/ollama"
  license "MIT"
  head "https://github.com/ollama/ollama.git", branch: "main"

  def install
    pkgshare.install "docs/crystals_package.md"
    pkgshare.install "api/types.go"
    pkgshare.install "server/routes.go"
    pkgshare.install "server/grax_consequence_contract.go"
  end

  caveats <<~EOS
    Crystals installs contract and routing source files only.
    It is intended to be consumed alongside the brullama CLI and grax contract package.
  EOS

  test do
    assert_predicate pkgshare/"docs/crystals_package.md", :exist?
    assert_predicate pkgshare/"server/routes.go", :exist?
  end
end
