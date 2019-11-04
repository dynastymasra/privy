CREATE TABLE IF NOT EXISTS categories (
    id bigserial PRIMARY KEY NOT NULL, -- IDENTIFIER
    name varchar(255) NOT NULL,
    enable bool NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id bigserial PRIMARY KEY NOT NULL, -- IDENTIFIER
    name varchar(255) NOT NULL,
    description text  NOT NULL,
    enable bool NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS images (
    id bigserial PRIMARY KEY NOT NULL, -- IDENTIFIER
    name varchar(255) NOT NULL,
    file text  NOT NULL,
    enable bool NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS category_products (
    product_id bigint  NOT NULL,
    category_id bigint NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE cascade ON UPDATE cascade,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE cascade ON UPDATE cascade
);

CREATE TABLE IF NOT EXISTS product_images (
    product_id bigint  NOT NULL,
    image_id bigint NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE cascade ON UPDATE cascade,
    FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE cascade ON UPDATE cascade
);

CREATE INDEX IF NOT EXISTS category_products_product_id_idx ON category_products (product_id);
CREATE INDEX IF NOT EXISTS category_products_category_id_idx ON category_products (category_id);
CREATE INDEX IF NOT EXISTS product_images_product_id_idx ON product_images (product_id);
CREATE INDEX IF NOT EXISTS product_images_image_id_idx ON product_images (image_id);