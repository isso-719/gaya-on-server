require 'bundler/setup'
Bundler.require
require 'sinatra/reloader' if development?
require './models.rb'


# ボタンの文字列集
def shapes
  ["heart", "good"]
end


# ブラウザ処理

get '/' do
  @logo = "images/logo-color.svg"
  
  erb :index
end

before '/room/:id' do
  if Room.find_by(room_id: params[:id]).nil?
    Room.create(
      room_id: params[:id]
    )
    shapes.each do |s|
      Count.create(
        room_id: Room.last.id,
        shape: s
      )
    end
  end
end

get '/room/:id' do
  @logo = "../images/logo-color.svg"
  
  @room = Room.find_by(room_id: params[:id])
  erb :room
end

post '/:id/:shape/inc' do
  count = Room.find_by(room_id: params[:id]).counts.find_by(shape: params[:shape])
  count.number += 1
  count.save
  redirect '/'
end





# electron処理

get '/:id/res' do
  begin
    count = Room.find_by(room_id: params[:id]).counts
    res = {}
    res.store("status", "success")
    
    temp = {}
    shapes.each do |s|
      temp.store(s, count.find_by(shape: s).number)
    end
    res.store("shapes", temp)

  rescue => e
    res = {"status": "failure" ,"error": e}
  end

  json res
end