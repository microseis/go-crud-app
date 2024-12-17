-- +goose Up
-- +goose StatementBegin
ALTER TABLE products ADD COLUMN user_id INT NOT NULL;
ALTER TABLE products
      ADD CONSTRAINT fk_user_products FOREIGN KEY (user_id) 
          REFERENCES users (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products DROP CONSTRAINT fk_user_products;
-- +goose StatementEnd
