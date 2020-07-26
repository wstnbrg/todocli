# This file was generated by GoReleaser. DO NOT EDIT.
class Todocli < Formula
  desc "ToDoCli is a simple cli saving your time by managing your tasks from your terminal."
  homepage ""
  version "1.0.3"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/wstnbrg/todocli/releases/download/1.0.3/todocli_1.0.3_Darwin_x86_64.tar.gz"
    sha256 "f1612ca78ab780b4215df9bec7f185d817a6810d1cede641d6751362696e9a09"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/wstnbrg/todocli/releases/download/1.0.3/todocli_1.0.3_Linux_x86_64.tar.gz"
      sha256 "1fefca45efb38ab3d358556621feff5764d1aa5dc2eb576cc3e94a97c6a8fad8"
    end
  end

  def install
    bin.install "todocli"
  end
end