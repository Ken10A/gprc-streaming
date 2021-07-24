# frozen_string_literal : true

$LOAD_PATH.unshift('.')
require 'proto.database_service_pb'

Stub    = Database::Stub
Request = SearchRequest

term ARGV.first

begin
    stub    = Stub.new('localhost:8080', :this_channel_is_insecure)
    request = Request.new(term: term, max_result: 10)

    stub.search(request).each do |reponse|
        puts "Term : #{response.matched_term}"
        puts "Rank : #{response.rank}"
        puts "Content: #{response.content}"
        puts
    end

rescure StarndardError => e
    code = e.respond_to?(:code) ? e.code : 'Unknown'
    puts "Code: #{code}, Type: '#{e.class}', Message: #{e.message}"
end
