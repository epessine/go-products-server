import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { CreateProductDto } from './dto/create-product.dto';
import { ClientGrpc } from '@nestjs/microservices';
import { ProductsClientGrpc } from './products.grpc';

@Injectable()
export class ProductsService implements OnModuleInit {
  private productsGrpcService: ProductsClientGrpc;

  constructor(
    @Inject('PRODUCTS_PACKAGE') private productsGrpcPackage: ClientGrpc,
  ) {}

  onModuleInit() {
    this.productsGrpcService =
      this.productsGrpcPackage.getService('ProductService');
  }

  create(createProductDto: CreateProductDto) {
    return this.productsGrpcService.createProduct(createProductDto);
  }

  findAll() {
    return this.productsGrpcService.findProducts({});
  }
}
