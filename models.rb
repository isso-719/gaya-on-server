ActiveRecord::Base.establish_connection
class Count < ActiveRecord::Base
    belongs_to :room

end

class Room < ActiveRecord::Base
    has_many :counts

end