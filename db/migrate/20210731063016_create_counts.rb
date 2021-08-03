class CreateCounts < ActiveRecord::Migration[6.1]
  def change
    create_table :counts do |t|
      t.references :room
      t.string :shape
      t.integer :number, default: 0
    end
  end
end
