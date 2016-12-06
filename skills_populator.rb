require 'rest-client'

skills = File.readlines("/Users/zach/development/squarewave_docker/api/src/skills_list.txt")
skills.map!{|s| s.strip}

skills.each_with_index do |skill, idx|
  puts "#{idx}/#{skills.length}: skill"
  RestClient.post("http://contactus.squarewaveng.com/skills", {name: skill}.to_json)
end
RestClient.post("http://contactus.squarewaveng.com/skills", {name: skills[0]}.to_json)
puts skills
