# This file was generated by GoReleaser. DO NOT EDIT.
class Todocli < Formula
  desc "ToDoCli is a simple cli saving your time by managing your tasks from your terminal."
  homepage ""
  version "1.0.5"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/wstnbrg/todocli/releases/download/1.0.5/todocli_1.0.5_Darwin_x86_64.tar.gz"
    sha256 "eeafadb5232722655137eb7f8be18ff9ffbb9175d07535de1cd81daaf4f0ddaf"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/wstnbrg/todocli/releases/download/1.0.5/todocli_1.0.5_Linux_x86_64.tar.gz"
      sha256 "4ca118cd77511e72ebd128f58c12d9350721f307660113de8b05aa1f6f271d7c"
    end
  end

  def install
    bin.install "todocli"
  end
end
