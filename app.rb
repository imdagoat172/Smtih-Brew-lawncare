require 'sinatra'
require 'json'

before do
  response.headers['Access-Control-Allow-Origin'] = '*'
  response.headers['Access-Control-Allow-Methods'] = 'GET, POST, OPTIONS'
  response.headers['Access-Control-Allow-Headers'] = 'Content-Type'
end

options '*' do
  200
end

quotes = []

get '/api/info' do
  content_type :json
  {
    name: 'Smith/Brew Lawncare LLC',
    type: 'Black-owned lawn care',
    services: ['lawn mowing','grass planting','edge trimming','leaf removal','bushes','flowers'],
    email: 'Marcussmith481@gmail.com',
    phone: '678-544-7911'
  }.to_json
end

get '/api/quotes' do
  content_type :json
  { count: quotes.size, quotes: quotes }.to_json
end

post '/api/quotes' do
  payload = JSON.parse(request.body.read) rescue {}
  required = %w[name email phone details]
  missing = required.select { |f| payload[f].to_s.strip.empty? }

  if missing.any?
    status 400
    return { error: "Missing fields: #{missing.join(', ')}" }.to_json
  end

  quote_id = quotes.size + 1
  quote = payload.merge('id' => quote_id, 'status' => 'pending')
  quotes << quote

  status 201
  { id: quote_id, message: 'Quote request received', quote: quote }.to_json
end
