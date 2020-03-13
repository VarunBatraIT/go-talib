require 'rest-client'
require 'oj'

ohlcvData = Array.new

ohlcTimeStamp = Array.new
ohlcOpen = Array.new
ohlcHigh = Array.new
ohlcLow = Array.new
ohlcClose = Array.new
ohlcVol = Array.new

json = RestClient.get("https://api.cryptowat.ch/markets/bitstamp/btcusd/ohlc?periods=86400")
v = Oj.load(json)

gct_ta = false

v['result']['86400'].each_with_index do |r, i|    
    next unless i > (v['result']['86400'].length() - 101)
    if gct_ta 
        puts "[#{r[0]}, #{r[1]}, #{r[2]}, #{r[3]}, #{r[4]}, #{r[5].round(2)}],"
    else
        ohlcTimeStamp.push(r[0])
        ohlcOpen.push(r[1])
        ohlcHigh.push(r[2])
        ohlcLow.push(r[3])
        ohlcClose.push(r[4])
        ohlcVol.push(r[5].round(2))
    end
end

puts "testTimestamp = []float64{#{ohlcTimeStamp.join(',')}}" 
puts "testOpen = []float64{#{ohlcOpen.join(',')}}" 
puts "testHigh = []float64{#{ohlcHigh.join(',')}}" 
puts "testLow = []float64{#{ohlcLow.join(',')}}" 
puts "testClose = []float64{#{ohlcClose.join(',')}}" 
puts "testVolume = []float64{#{ohlcVol.join(',')}}" 