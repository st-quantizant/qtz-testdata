# rake: Ruby build automation tool — no cryptography
require 'fileutils'

module Rake
  class Task
    attr_reader :name, :dependencies

    def initialize(name, &block)
      @name = name
      @dependencies = []
      @block = block
    end

    def invoke
      @dependencies.each { |dep| TaskManager.find(dep)&.invoke }
      @block&.call
    end
  end

  module TaskManager
    @tasks = {}

    def self.define(name, deps = [], &block)
      t = Task.new(name, &block)
      t.dependencies.concat(deps)
      @tasks[name.to_s] = t
    end

    def self.find(name)
      @tasks[name.to_s]
    end
  end

  def self.task(name, deps = [], &block)
    TaskManager.define(name, deps, &block)
  end
end
