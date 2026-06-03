# json gem: JSON parsing and generation — no cryptography
require 'json'

module JSONHelper
  def self.parse(str)
    JSON.parse(str)
  end

  def self.dump(obj)
    JSON.generate(obj)
  end

  def self.pretty(obj)
    JSON.pretty_generate(obj)
  end
end
