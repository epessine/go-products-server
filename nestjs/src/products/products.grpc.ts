import { Observable } from 'rxjs';

interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
}

export interface CreateProductRequest {
  name: string;
  description: string;
  price: number;
}

export interface FindProductsRequest {}

export interface CreateProductResponse {
  product: Product;
}

export interface FindProductsResponse {
  products: Product[];
}

export interface ProductsClientGrpc {
  createProduct: (
    data: CreateProductRequest,
  ) => Observable<CreateProductResponse>;
  findProducts: (data: FindProductsRequest) => Observable<FindProductsResponse>;
}
