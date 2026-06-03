# activesupport: Rails utility library — no cryptography primitives
require 'time'
require 'date'

module ActiveSupport
  module CoreExtensions
    module String
      def camelize
        split('_').map(&:capitalize).join
      end

      def underscore
        gsub(/([A-Z]+)([A-Z][a-z])/, '\1_\2')
          .gsub(/([a-z\d])([A-Z])/, '\1_\2')
          .downcase
      end

      def pluralize
        return self if end_with?('s')
        end_with?('y') ? self[0..-2] + 'ies' : self + 's'
      end
    end

    module Array
      def flatten_once
        inject([]) { |acc, el| el.is_a?(::Array) ? acc.concat(el) : acc << el }
      end
    end
  end

  def self.now
    Time.now.utc
  end
end

class String; include ActiveSupport::CoreExtensions::String; end
class Array;  include ActiveSupport::CoreExtensions::Array;  end
